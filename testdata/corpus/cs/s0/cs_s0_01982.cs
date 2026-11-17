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

    public delegate void Notifier18470(int value);

    public sealed class Source18470
    {
        public event Notifier18470 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate18471(IEnumerable<int> source)
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
    public struct Point18472
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid18472
    {
        public static Point18472 Make(int left, int right)
        {
            return new Point18472 { X = left, Y = right };
        }
    }
    public static int Scale18473<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public static int Scale18474<T>(T value) where T : struct
    {
        int number = 85;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public static int Scale18475<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public sealed class SlotMap18476
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
}
