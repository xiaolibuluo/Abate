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

    public struct Point18410
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid18410
    {
        public static Point18410 Make(int left, int right)
        {
            return new Point18410 { X = left, Y = right };
        }
    }
    public delegate void Notifier18411(int value);

    public sealed class Source18411
    {
        public event Notifier18411 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate18412(IEnumerable<int> source)
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
    public static int Scale18413<T>(T value) where T : struct
    {
        int number = 97;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 58; }
        return number;
    }
    public sealed class SlotMap18414
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
