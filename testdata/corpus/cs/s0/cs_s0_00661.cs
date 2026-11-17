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

    public sealed class SlotMap6160
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
    public static int Scale6161<T>(T value) where T : struct
    {
        int number = 66;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 62; }
        return number;
    }
    public struct Point6162
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 73;
    }

    public static class Grid6162
    {
        public static Point6162 Make(int left, int right)
        {
            return new Point6162 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate6163(IEnumerable<int> source)
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
}
