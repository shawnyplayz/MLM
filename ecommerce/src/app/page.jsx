'use client';

import Link from 'next/link';
import { ShoppingBag, Heart, Award, Leaf, Shield, Zap, Users, TrendingUp } from 'lucide-react';

export default function Home() {
  const featuredProducts = [
    {
      id: 1,
      name: 'Premium Multivitamin Complex',
      price: 49.99,
      image: 'https://images.unsplash.com/photo-1584308666744-24d5c474f2ae?w=400&h=400&fit=crop',
      category: 'Vitamins',
      description: 'Complete daily nutrition in one capsule'
    },
    {
      id: 2,
      name: 'Omega-3 Fish Oil',
      price: 34.99,
      image: 'https://images.unsplash.com/photo-1505751172876-fa1923c5c528?w=400&h=400&fit=crop',
      category: 'Essential Fatty Acids',
      description: 'Pure EPA & DHA for heart health'
    },
    {
      id: 3,
      name: 'Protein Power Plus',
      price: 59.99,
      image: 'https://images.unsplash.com/photo-1579722821273-0f6c7d44362f?w=400&h=400&fit=crop',
      category: 'Protein',
      description: '25g protein per serving'
    },
    {
      id: 4,
      name: 'Green Superfood Blend',
      price: 44.99,
      image: 'https://images.unsplash.com/photo-1607623814075-e51df1bdc82f?w=400&h=400&fit=crop',
      category: 'Superfoods',
      description: 'Organic greens & antioxidants'
    }
  ];

  return (
    <div className="min-h-screen bg-white">
      {/* Header */}
      <header className="border-b bg-white shadow-sm sticky top-0 z-50">
        <div className="container mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <Link href="/" className="flex items-center space-x-2">
              <Leaf className="h-8 w-8 text-green-600" />
              <span className="text-2xl font-bold text-gray-900">VitaLife</span>
            </Link>
            
            <nav className="hidden md:flex space-x-8">
              <Link href="/products" className="text-gray-700 hover:text-green-600 font-medium">Shop</Link>
              <Link href="/categories" className="text-gray-700 hover:text-green-600 font-medium">Categories</Link>
              <Link href="/opportunity" className="text-gray-700 hover:text-green-600 font-medium">Business Opportunity</Link>
              <Link href="/about" className="text-gray-700 hover:text-green-600 font-medium">About</Link>
            </nav>
            
            <div className="flex items-center space-x-4">
              <button className="p-2 hover:bg-gray-100 rounded-full">
                <ShoppingBag className="h-6 w-6 text-gray-700" />
              </button>
              <Link href="/login" className="px-4 py-2 text-green-600 hover:text-green-700 font-medium">
                Login
              </Link>
              <Link href="/register" className="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 font-medium">
                Join Now
              </Link>
            </div>
          </div>
        </div>
      </header>

      {/* Hero Section */}
      <section className="bg-gradient-to-br from-green-50 via-white to-blue-50 py-20">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-2 gap-12 items-center">
            <div>
              <div className="inline-block px-4 py-2 bg-green-100 text-green-700 rounded-full text-sm font-semibold mb-4">
                Premium Food Supplements
              </div>
              <h1 className="text-5xl md:text-6xl font-bold text-gray-900 mb-6">
                Fuel Your Body,<br />
                <span className="text-green-600">Build Your Future</span>
              </h1>
              <p className="text-xl text-gray-600 mb-8">
                Premium quality supplements that transform health. Join our community and earn while helping others achieve wellness.
              </p>
              <div className="flex flex-col sm:flex-row gap-4">
                <Link href="/products" className="px-8 py-4 bg-green-600 text-white rounded-lg text-lg font-semibold hover:bg-green-700 transition text-center">
                  Shop Supplements
                </Link>
                <Link href="/opportunity" className="px-8 py-4 bg-white text-green-600 border-2 border-green-600 rounded-lg text-lg font-semibold hover:bg-green-50 transition text-center">
                  Business Opportunity
                </Link>
              </div>
              <div className="mt-8 flex items-center space-x-8">
                <div>
                  <div className="text-3xl font-bold text-gray-900">10,000+</div>
                  <div className="text-sm text-gray-600">Happy Customers</div>
                </div>
                <div>
                  <div className="text-3xl font-bold text-gray-900">5,000+</div>
                  <div className="text-sm text-gray-600">Distributors</div>
                </div>
                <div>
                  <div className="text-3xl font-bold text-gray-900">100%</div>
                  <div className="text-sm text-gray-600">Natural</div>
                </div>
              </div>
            </div>
            <div className="relative">
              <div className="bg-gradient-to-br from-green-400 to-blue-500 rounded-3xl p-8 shadow-2xl">
                <div className="bg-white rounded-2xl p-6">
                  <img src="https://images.unsplash.com/photo-1556911220-bff31c812dba?w=800&h=600&fit=crop" alt="Premium Supplements" className="w-full h-96 object-cover rounded-xl" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Trust Badges */}
      <section className="py-12 bg-gray-50">
        <div className="container mx-auto px-4">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            <div className="flex flex-col items-center text-center">
              <Shield className="h-12 w-12 text-green-600 mb-2" />
              <h3 className="font-semibold text-gray-900">FDA Approved</h3>
              <p className="text-sm text-gray-600">Certified Facility</p>
            </div>
            <div className="flex flex-col items-center text-center">
              <Leaf className="h-12 w-12 text-green-600 mb-2" />
              <h3 className="font-semibold text-gray-900">100% Natural</h3>
              <p className="text-sm text-gray-600">No Artificial Additives</p>
            </div>
            <div className="flex flex-col items-center text-center">
              <Heart className="h-12 w-12 text-green-600 mb-2" />
              <h3 className="font-semibold text-gray-900">Science-Backed</h3>
              <p className="text-sm text-gray-600">Clinically Tested</p>
            </div>
            <div className="flex flex-col items-center text-center">
              <Award className="h-12 w-12 text-green-600 mb-2" />
              <h3 className="font-semibold text-gray-900">Award Winning</h3>
              <p className="text-sm text-gray-600">Industry Leader</p>
            </div>
          </div>
        </div>
      </section>

      {/* Featured Products */}
      <section className="py-20">
        <div className="container mx-auto px-4">
          <div className="text-center mb-12">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Featured Supplements</h2>
            <p className="text-xl text-gray-600">Premium quality products for optimal health</p>
          </div>
          
          <div className="grid md:grid-cols-4 gap-8">
            {featuredProducts.map((product) => (
              <div key={product.id} className="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-xl transition group">
                <div className="relative h-64 bg-gray-100">
                  <img src={product.image} alt={product.name} className="w-full h-full object-cover" />
                  <div className="absolute top-4 right-4 bg-green-600 text-white px-3 py-1 rounded-full text-sm font-semibold">
                    New
                  </div>
                </div>
                <div className="p-6">
                  <div className="text-sm text-green-600 font-semibold mb-2">{product.category}</div>
                  <h3 className="text-lg font-bold text-gray-900 mb-2">{product.name}</h3>
                  <p className="text-sm text-gray-600 mb-4">{product.description}</p>
                  <div className="flex items-center justify-between">
                    <span className="text-2xl font-bold text-gray-900">${product.price}</span>
                    <button className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition">
                      Add to Cart
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
          
          <div className="text-center mt-12">
            <Link href="/products" className="inline-block px-8 py-4 bg-gray-900 text-white rounded-lg text-lg font-semibold hover:bg-gray-800 transition">
              View All Products
            </Link>
          </div>
        </div>
      </section>

      {/* Categories */}
      <section className="py-20 bg-gray-50">
        <div className="container mx-auto px-4">
          <div className="text-center mb-12">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Shop by Category</h2>
            <p className="text-xl text-gray-600">Find the perfect supplement for your needs</p>
          </div>
          
          <div className="grid md:grid-cols-3 gap-8">
            {[
              { name: 'Vitamins & Minerals', count: 45, icon: '💊' },
              { name: 'Protein & Fitness', count: 32, icon: '💪' },
              { name: 'Herbal Supplements', count: 28, icon: '🌿' },
              { name: 'Immune Support', count: 24, icon: '🛡️' },
              { name: 'Digestive Health', count: 18, icon: '🫁' },
              { name: 'Energy & Focus', count: 21, icon: '⚡' }
            ].map((category, index) => (
              <Link key={index} href={`/category/${category.name.toLowerCase().replace(/ /g, '-')}`} className="bg-white p-8 rounded-xl shadow-lg hover:shadow-xl transition group">
                <div className="text-5xl mb-4">{category.icon}</div>
                <h3 className="text-xl font-bold text-gray-900 mb-2 group-hover:text-green-600 transition">{category.name}</h3>
                <p className="text-gray-600">{category.count} Products</p>
              </Link>
            ))}
          </div>
        </div>
      </section>

      {/* MLM Opportunity */}
      <section className="py-20 bg-gradient-to-br from-green-600 to-blue-600 text-white">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-2 gap-12 items-center">
            <div>
              <h2 className="text-4xl md:text-5xl font-bold mb-6">
                Build Your Wellness Empire
              </h2>
              <p className="text-xl text-green-50 mb-8">
                Join thousands of successful distributors earning income while helping others achieve better health. Our proven MLM system offers multiple income streams and unlimited growth potential.
              </p>
              <div className="grid grid-cols-2 gap-6 mb-8">
                <div className="bg-white/10 backdrop-blur-sm rounded-lg p-6">
                  <Users className="h-10 w-10 mb-3" />
                  <h3 className="font-bold text-lg mb-2">Build Your Team</h3>
                  <p className="text-sm text-green-50">Multiple tree structures supported</p>
                </div>
                <div className="bg-white/10 backdrop-blur-sm rounded-lg p-6">
                  <TrendingUp className="h-10 w-10 mb-3" />
                  <h3 className="font-bold text-lg mb-2">Earn Commissions</h3>
                  <p className="text-sm text-green-50">Up to 20% on sales</p>
                </div>
              </div>
              <Link href="/opportunity" className="inline-block px-8 py-4 bg-white text-green-600 rounded-lg text-lg font-semibold hover:bg-gray-100 transition">
                Learn More About Our Opportunity
              </Link>
            </div>
            <div className="bg-white/10 backdrop-blur-sm rounded-2xl p-8">
              <h3 className="text-2xl font-bold mb-6">Income Potential</h3>
              <div className="space-y-4">
                {[
                  { rank: 'Bronze', income: '$500-$1,000/mo', members: '5+' },
                  { rank: 'Silver', income: '$1,000-$3,000/mo', members: '15+' },
                  { rank: 'Gold', income: '$3,000-$7,000/mo', members: '30+' },
                  { rank: 'Platinum', income: '$7,000-$15,000/mo', members: '50+' },
                  { rank: 'Diamond', income: '$15,000+/mo', members: '100+' }
                ].map((tier, index) => (
                  <div key={index} className="flex items-center justify-between bg-white/10 rounded-lg p-4">
                    <div>
                      <div className="font-bold">{tier.rank}</div>
                      <div className="text-sm text-green-50">{tier.members} team members</div>
                    </div>
                    <div className="text-xl font-bold">{tier.income}</div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Testimonials */}
      <section className="py-20">
        <div className="container mx-auto px-4">
          <div className="text-center mb-12">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">What Our Customers Say</h2>
            <p className="text-xl text-gray-600">Real results from real people</p>
          </div>
          
          <div className="grid md:grid-cols-3 gap-8">
            {[
              {
                name: 'Sarah Johnson',
                role: 'Gold Distributor',
                text: 'These supplements changed my life! And the business opportunity has given me financial freedom.',
                rating: 5
              },
              {
                name: 'Michael Chen',
                role: 'Platinum Distributor',
                text: 'Best decision I ever made. Quality products and an amazing income opportunity.',
                rating: 5
              },
              {
                name: 'Emily Rodriguez',
                role: 'Diamond Distributor',
                text: 'From customer to top earner in 2 years. The products sell themselves!',
                rating: 5
              }
            ].map((testimonial, index) => (
              <div key={index} className="bg-white p-8 rounded-xl shadow-lg">
                <div className="flex mb-4">
                  {[...Array(testimonial.rating)].map((_, i) => (
                    <span key={i} className="text-yellow-400 text-xl">★</span>
                  ))}
                </div>
                <p className="text-gray-700 mb-6 italic">"{testimonial.text}"</p>
                <div className="flex items-center">
                  <div className="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center text-green-600 font-bold text-xl mr-4">
                    {testimonial.name[0]}
                  </div>
                  <div>
                    <div className="font-bold text-gray-900">{testimonial.name}</div>
                    <div className="text-sm text-gray-600">{testimonial.role}</div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 bg-gray-900 text-white">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Ready to Transform Your Health & Wealth?
          </h2>
          <p className="text-xl text-gray-300 mb-8 max-w-2xl mx-auto">
            Join VitaLife today and start your journey to better health and financial freedom
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link href="/products" className="px-8 py-4 bg-green-600 text-white rounded-lg text-lg font-semibold hover:bg-green-700 transition">
              Shop Now
            </Link>
            <Link href="/register" className="px-8 py-4 bg-white text-gray-900 rounded-lg text-lg font-semibold hover:bg-gray-100 transition">
              Become a Distributor
            </Link>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-gray-400 py-12 border-t border-gray-800">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-5 gap-8">
            <div className="md:col-span-2">
              <div className="flex items-center space-x-2 mb-4">
                <Leaf className="h-8 w-8 text-green-600" />
                <span className="text-2xl font-bold text-white">VitaLife</span>
              </div>
              <p className="text-sm mb-4">Premium food supplements for optimal health and wellness. Join our MLM network and build your future.</p>
              <div className="flex space-x-4">
                <a href="#" className="hover:text-white">Facebook</a>
                <a href="#" className="hover:text-white">Instagram</a>
                <a href="#" className="hover:text-white">Twitter</a>
              </div>
            </div>
            <div>
              <h4 className="text-white font-bold mb-4">Shop</h4>
              <ul className="space-y-2 text-sm">
                <li><Link href="/products" className="hover:text-white">All Products</Link></li>
                <li><Link href="/categories" className="hover:text-white">Categories</Link></li>
                <li><Link href="/bestsellers" className="hover:text-white">Best Sellers</Link></li>
                <li><Link href="/new" className="hover:text-white">New Arrivals</Link></li>
              </ul>
            </div>
            <div>
              <h4 className="text-white font-bold mb-4">Business</h4>
              <ul className="space-y-2 text-sm">
                <li><Link href="/opportunity" className="hover:text-white">Join Us</Link></li>
                <li><Link href="/compensation" className="hover:text-white">Compensation Plan</Link></li>
                <li><Link href="/training" className="hover:text-white">Training</Link></li>
                <li><Link href="/success-stories" className="hover:text-white">Success Stories</Link></li>
              </ul>
            </div>
            <div>
              <h4 className="text-white font-bold mb-4">Support</h4>
              <ul className="space-y-2 text-sm">
                <li><Link href="/contact" className="hover:text-white">Contact Us</Link></li>
                <li><Link href="/faq" className="hover:text-white">FAQ</Link></li>
                <li><Link href="/shipping" className="hover:text-white">Shipping</Link></li>
                <li><Link href="/returns" className="hover:text-white">Returns</Link></li>
              </ul>
            </div>
          </div>
          <div className="border-t border-gray-800 mt-8 pt-8 text-center text-sm">
            <p>&copy; 2024 VitaLife. All rights reserved. | <Link href="/privacy" className="hover:text-white">Privacy Policy</Link> | <Link href="/terms" className="hover:text-white">Terms of Service</Link></p>
          </div>
        </div>
      </footer>
    </div>
  );
}
