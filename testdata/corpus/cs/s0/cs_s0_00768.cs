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

    public static int Scale7200<T>(T value) where T : struct
    {
        int number = 53;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public sealed class SlotMap7201
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
    public struct Point7202
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 67;
    }

    public static class Grid7202
    {
        public static Point7202 Make(int left, int right)
        {
            return new Point7202 { X = left, Y = right };
        }
    }
    public static int Scale7203<T>(T value) where T : struct
    {
        int number = 74;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
    public async Task<int> Aggregate7204(IEnumerable<int> source)
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
