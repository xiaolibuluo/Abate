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

    public async Task<int> Aggregate12890(IEnumerable<int> source)
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
    public struct Point12891
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 85;
    }

    public static class Grid12891
    {
        public static Point12891 Make(int left, int right)
        {
            return new Point12891 { X = left, Y = right };
        }
    }
    public static int Scale12892<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 26; }
        return number;
    }
}
