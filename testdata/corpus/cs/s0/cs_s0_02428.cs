using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public struct Point22680
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid22680
    {
        public static Point22680 Make(int left, int right)
        {
            return new Point22680 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate22681(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public struct Point22682
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid22682
    {
        public static Point22682 Make(int left, int right)
        {
            return new Point22682 { X = left, Y = right };
        }
    }
}
