import { useAppSelector } from '../../store/hooks';

export default function DistributorDashboard() {
  const { user } = useAppSelector((state) => state.auth);

  return (
    <div className="min-h-screen bg-gray-100">
      <nav className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 py-4">
          <h1 className="text-2xl font-bold">My Dashboard</h1>
        </div>
      </nav>
      
      <div className="max-w-7xl mx-auto px-4 py-8">
        <div className="bg-white rounded-lg shadow p-6 mb-6">
          <h2 className="text-xl font-semibold mb-4">Welcome, {user?.first_name}!</h2>
          <p className="text-gray-600">Role: Distributor</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-gray-500">My Sales</h3>
            <p className="text-3xl font-bold mt-2">$0</p>
          </div>
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-gray-500">My Commissions</h3>
            <p className="text-3xl font-bold mt-2">$0</p>
          </div>
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-gray-500">Team Members</h3>
            <p className="text-3xl font-bold mt-2">0</p>
          </div>
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-gray-500">My Rank</h3>
            <p className="text-3xl font-bold mt-2">Bronze</p>
          </div>
        </div>

        <div className="mt-8 bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-semibold mb-4">Distributor Features</h3>
          <ul className="space-y-2">
            <li>• View my team/downline</li>
            <li>• Track my commissions</li>
            <li>• Add new team members</li>
            <li>• View my reports</li>
            <li>• Manage my profile</li>
          </ul>
        </div>
      </div>
    </div>
  );
}
