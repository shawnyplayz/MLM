'use client';

import { useState } from 'react';
import Link from 'next/link';
import { ShoppingBag, Leaf, Filter, Search } from 'lucide-react';

export default function ProductsPage() {
  const [selectedCategory, setSelectedCategory] = useState('all');
  const [searchQuery, setSearchQuery] = useState('');

  const categories = [
    { id: 'all', name: 'All Products', count: 156 },
    { id: 'vitamins', name: 'Vitamins & Minerals', count: 45 },
    { id: 'protein', name: 'Protein & Fitness', count: 32 },
    { id: 'herbal', name: 'Herbal Supplements', count: 28 },
    { id: 'immune', name: 'Immune Support', count: 24 },
    { id: 'digestive', name: 'Digestive Health', count: 18 },
    { id: 'energy', name: 'Energy & Focus', count: 21 }
  ];

  const products = [
    {
      id: 1,
      name: 'Premium Multivitamin Complex',
      price: 49.99,
      category: 'vitamins',
      rating: 4.8,
      reviews: 234,
      inStock: true,
      badge: 'Best Seller',
      image: 'https://images.unsplash.com/photo-1584308666744-24d5c474f2ae?w=400&h=400&fit=crop',
      description: 'Complete daily nutrition with 25+ essential vitamins and minerals'
    },
    {
      id: 2,
      name: 'Omega-3 Fish Oil 1000mg',
      price: 34.99,
      category: 'vitamins',
      rating: 4.9,
      reviews: 189,
      inStock: true,
      badge: 'Top Rated',
      image: 'https://images.unsplash.com/photo-1505751172876-fa1923c5c528?w=400&h=400&fit=crop',
      description: 'Pure EPA & DHA for heart and brain health'
    },
    {
      id: 3,
      name: 'Whey Protein Isolate',
      price: 59.99,
      category: 'protein',
      rating: 4.7,
      reviews: 456,
      inStock: true,
      badge: 'New',
      image: 'https://images.unsplash.com/photo-1579722821273-0f6c7d44362f?w=400&h=400&fit=crop',
      description: '25g protein per serving, zero sugar'
    },
    {
      id: 4,
      name: 'Green Superfood Powder',
      price: 44.99,
      category: 'herbal',
      rating: 4.6,
      reviews: 178,
      inStock: true,
      badge: null,
      image: 'https://images.unsplash.com/photo-1607623814075-e51df1bdc82f?w=400&h=400&fit=crop',
      description: 'Organic greens, fruits & probiotics'
    },
    {
      id: 5,
      name: 'Vitamin D3 5000 IU',
      price: 24.99,
      category: 'vitamins',
      rating: 4.9,
      reviews: 312,
      inStock: true,
      badge: 'Best Seller',
      image: 'https://images.unsplash.com/photo-1550572017-edd951aa8f72?w=400&h=400&fit=crop',
      description: 'Supports bone health and immune function'
    },
    {
      id: 6,
      name: 'Turmeric Curcumin Complex',
      price: 39.99,
      category: 'herbal',
      rating: 4.8,
      reviews: 267,
      inStock: true,
      badge: null,
      image: 'https://images.unsplash.com/photo-1615485500704-8e990f9900f7?w=400&h=400&fit=crop',
      description: 'Anti-inflammatory support with BioPerine'
    },
    {
      id: 7,
      name: 'Probiotic 50 Billion CFU',
      price: 49.99,
      category: 'digestive',
      rating: 4.7,
      reviews: 198,
      inStock: true,
      badge: 'Top Rated',
      image: 'https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=400&h=400&fit=crop',
      description: '10 strains for digestive & immune health'
    },
    {
      id: 8,
      name: 'Collagen Peptides',
      price: 54.99,
      category: 'protein',
      rating: 4.8,
      reviews: 289,
      inStock: true,
      badge: 'New',
      image: 'https://images.unsplash.com/photo-1608571423902-eed4a5ad8108?w=400&h=400&fit=crop',
      description: 'Type I & III collagen for skin, hair & joints'
    },
    {
      id: 9,
      name: 'Immune Defense Complex',
      price: 42.99,
      category: 'immune',
      rating: 4.9,
      reviews: 345,
      inStock: true,
      badge: 'Best Seller',
      image: 'https://images.unsplash.com/photo-1471864190281-a93a3070b6de?w=400&h=400&fit=crop',
      description: 'Vitamin C, Zinc, Elderberry & Echinacea'
    },
    {
      id: 10,
      name: 'Energy & Focus Formula',
      price: 37.99,
      category: 'energy',
      rating: 4.6,
      reviews: 156,
      inStock: true,
      badge: null,
      image: 'https://images.unsplash.com/photo-1610348725531-843dff563e2c?w=400&h=400&fit=crop',
      description: 'Natural caffeine, L-theanine & B-vitamins'
    },
    {
      id: 11,
      name: 'Magnesium Glycinate 400mg',
      price: 29.99,
      category: 'vitamins',
      rating: 4.8,
      reviews: 223,
      inStock: true,
      badge: null,
      image: 'https://images.unsplash.com/photo-1526336024174-e58f5cdd8e13?w=400&h=400&fit=crop',
      description: 'Highly absorbable for sleep & relaxation'
    },
    {
      id: 12,
      name: 'Ashwagandha Root Extract',
      price: 34.99,
      category: 'herbal',
      rating: 4.7,
      reviews: 201,
      inStock: true,
      badge: 'Top Rated',
      image: 'https://images.unsplash.com/photo-1599932550619-d4e7c0d5d430?w=400&h=400&fit=crop',
      description: 'Adaptogen for stress relief & energy'
    }
  ];

  const filteredProducts = products.filter(product => {
    const matchesCategory = selectedCategory === 'all' || product.category === selectedCategory;
    const matchesSearch = product.name.toLowerCase().includes(searchQuery.toLowerCase());
    return matchesCategory && matchesSearch;
  });

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="border-b bg-white shadow-sm sticky top-0 z-50">
        <div className="container mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <Link href="/" className="flex items-center space-x-2">
              <Leaf className="h-8 w-8 text-green-600" />
              <span className="text-2xl font-bold text-gray-900">VitaLife</span>
            </Link>
            
            <nav className="hidden md:flex space-x-8">
              <Link href="/products" className="text-green-600 font-medium">Shop</Link>
              <Link href="/categories" className="text-gray-700 hover:text-green-600 font-medium">Categories</Link>
              <Link href="/opportunity" className="text-gray-700 hover:text-green-600 font-medium">Business Opportunity</Link>
              <Link href="/about" className="text-gray-700 hover:text-green-600 font-medium">About</Link>
            </nav>
            
            <div className="flex items-center space-x-4">
              <button className="p-2 hover:bg-gray-100 rounded-full relative">
                <ShoppingBag className="h-6 w-6 text-gray-700" />
                <span className="absolute -top-1 -right-1 bg-green-600 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">0</span>
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

      {/* Page Header */}
      <div className="bg-gradient-to-r from-green-600 to-blue-600 text-white py-12">
        <div className="container mx-auto px-4">
          <h1 className="text-4xl font-bold mb-2">Premium Supplements</h1>
          <p className="text-green-50">Discover our complete range of high-quality food supplements</p>
        </div>
      </div>

      <div className="container mx-auto px-4 py-8">
        <div className="grid md:grid-cols-4 gap-8">
          {/* Sidebar */}
          <div className="md:col-span-1">
            <div className="bg-white rounded-lg shadow-sm p-6 sticky top-24">
              <div className="flex items-center justify-between mb-4">
                <h2 className="text-lg font-bold text-gray-900">Filters</h2>
                <Filter className="h-5 w-5 text-gray-500" />
              </div>
              
              {/* Search */}
              <div className="mb-6">
                <label className="block text-sm font-medium text-gray-700 mb-2">Search</label>
                <div className="relative">
                  <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
                  <input
                    type="text"
                    placeholder="Search products..."
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  />
                </div>
              </div>

              {/* Categories */}
              <div>
                <h3 className="text-sm font-medium text-gray-700 mb-3">Categories</h3>
                <div className="space-y-2">
                  {categories.map((category) => (
                    <button
                      key={category.id}
                      onClick={() => setSelectedCategory(category.id)}
                      className={`w-full text-left px-3 py-2 rounded-lg transition ${
                        selectedCategory === category.id
                          ? 'bg-green-50 text-green-700 font-medium'
                          : 'text-gray-700 hover:bg-gray-50'
                      }`}
                    >
                      <div className="flex items-center justify-between">
                        <span>{category.name}</span>
                        <span className="text-sm text-gray-500">{category.count}</span>
                      </div>
                    </button>
                  ))}
                </div>
              </div>

              {/* Price Range */}
              <div className="mt-6 pt-6 border-t">
                <h3 className="text-sm font-medium text-gray-700 mb-3">Price Range</h3>
                <div className="space-y-2">
                  <label className="flex items-center">
                    <input type="checkbox" className="rounded text-green-600 mr-2" />
                    <span className="text-sm text-gray-700">Under $25</span>
                  </label>
                  <label className="flex items-center">
                    <input type="checkbox" className="rounded text-green-600 mr-2" />
                    <span className="text-sm text-gray-700">$25 - $50</span>
                  </label>
                  <label className="flex items-center">
                    <input type="checkbox" className="rounded text-green-600 mr-2" />
                    <span className="text-sm text-gray-700">$50 - $75</span>
                  </label>
                  <label className="flex items-center">
                    <input type="checkbox" className="rounded text-green-600 mr-2" />
                    <span className="text-sm text-gray-700">Over $75</span>
                  </label>
                </div>
              </div>
            </div>
          </div>

          {/* Products Grid */}
          <div className="md:col-span-3">
            <div className="flex items-center justify-between mb-6">
              <p className="text-gray-600">
                Showing <span className="font-semibold">{filteredProducts.length}</span> products
              </p>
              <select className="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent">
                <option>Sort by: Featured</option>
                <option>Price: Low to High</option>
                <option>Price: High to Low</option>
                <option>Name: A to Z</option>
                <option>Best Rating</option>
              </select>
            </div>

            <div className="grid md:grid-cols-3 gap-6">
              {filteredProducts.map((product) => (
                <div key={product.id} className="bg-white rounded-lg shadow-sm overflow-hidden hover:shadow-lg transition group">
                  <div className="relative h-64 bg-gray-100">
                    <img src={product.image} alt={product.name} className="w-full h-full object-cover" />
                    {product.badge && (
                      <div className="absolute top-4 right-4 bg-green-600 text-white px-3 py-1 rounded-full text-xs font-semibold">
                        {product.badge}
                      </div>
                    )}
                    {!product.inStock && (
                      <div className="absolute inset-0 bg-black/50 flex items-center justify-center">
                        <span className="bg-red-600 text-white px-4 py-2 rounded-lg font-semibold">Out of Stock</span>
                      </div>
                    )}
                  </div>
                  
                  <div className="p-5">
                    <div className="flex items-center mb-2">
                      <div className="flex items-center">
                        {[...Array(5)].map((_, i) => (
                          <span key={i} className={`text-sm ${i < Math.floor(product.rating) ? 'text-yellow-400' : 'text-gray-300'}`}>
                            ★
                          </span>
                        ))}
                      </div>
                      <span className="text-sm text-gray-600 ml-2">({product.reviews})</span>
                    </div>
                    
                    <h3 className="text-lg font-bold text-gray-900 mb-2 group-hover:text-green-600 transition">
                      {product.name}
                    </h3>
                    
                    <p className="text-sm text-gray-600 mb-4 line-clamp-2">
                      {product.description}
                    </p>
                    
                    <div className="flex items-center justify-between">
                      <span className="text-2xl font-bold text-gray-900">${product.price}</span>
                      <button 
                        disabled={!product.inStock}
                        className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition disabled:bg-gray-300 disabled:cursor-not-allowed"
                      >
                        Add to Cart
                      </button>
                    </div>
                    
                    <Link href={`/products/${product.id}`} className="block mt-3 text-sm text-green-600 hover:text-green-700 font-medium">
                      View Details →
                    </Link>
                  </div>
                </div>
              ))}
            </div>

            {filteredProducts.length === 0 && (
              <div className="text-center py-12">
                <p className="text-gray-600 text-lg">No products found matching your criteria.</p>
                <button
                  onClick={() => {
                    setSelectedCategory('all');
                    setSearchQuery('');
                  }}
                  className="mt-4 px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
                >
                  Clear Filters
                </button>
              </div>
            )}
          </div>
        </div>
      </div>

      {/* MLM CTA */}
      <section className="py-16 bg-gradient-to-r from-green-600 to-blue-600 text-white">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-3xl font-bold mb-4">Love Our Products? Earn While You Share!</h2>
          <p className="text-xl text-green-50 mb-6">
            Become a VitaLife distributor and earn commissions on every sale
          </p>
          <Link href="/opportunity" className="inline-block px-8 py-4 bg-white text-green-600 rounded-lg text-lg font-semibold hover:bg-gray-100 transition">
            Learn About Our Business Opportunity
          </Link>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-gray-400 py-12 border-t border-gray-800">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-4 gap-8">
            <div>
              <div className="flex items-center space-x-2 mb-4">
                <Leaf className="h-8 w-8 text-green-600" />
                <span className="text-2xl font-bold text-white">VitaLife</span>
              </div>
              <p className="text-sm">Premium food supplements for optimal health</p>
            </div>
            <div>
              <h4 className="text-white font-bold mb-4">Shop</h4>
              <ul className="space-y-2 text-sm">
                <li><Link href="/products" className="hover:text-white">All Products</Link></li>
                <li><Link href="/bestsellers" className="hover:text-white">Best Sellers</Link></li>
                <li><Link href="/new" className="hover:text-white">New Arrivals</Link></li>
              </ul>
            </div>
            <div>
              <h4 className="text-white font-bold mb-4">Business</h4>
              <ul className="space-y-2 text-sm">
                <li><Link href="/opportunity" className="hover:text-white">Join Us</Link></li>
                <li><Link href="/compensation" className="hover:text-white">Compensation</Link></li>
                <li><Link href="/training" className="hover:text-white">Training</Link></li>
              </ul>
            </div>
            <div>
              <h4 className="text-white font-bold mb-4">Support</h4>
              <ul className="space-y-2 text-sm">
                <li><Link href="/contact" className="hover:text-white">Contact</Link></li>
                <li><Link href="/faq" className="hover:text-white">FAQ</Link></li>
                <li><Link href="/shipping" className="hover:text-white">Shipping</Link></li>
              </ul>
            </div>
          </div>
          <div className="border-t border-gray-800 mt-8 pt-8 text-center text-sm">
            <p>&copy; 2024 VitaLife. All rights reserved.</p>
          </div>
        </div>
      </footer>
    </div>
  );
}
