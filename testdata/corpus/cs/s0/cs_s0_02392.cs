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

    public delegate void Notifier22340(int value);

    public sealed class Source22340
    {
        public event Notifier22340 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22341<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
    public sealed class SlotMap22342
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
    public async Task<int> Aggregate22343(IEnumerable<int> source)
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
    public struct Point22344
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 72;
    }

    public static class Grid22344
    {
        public static Point22344 Make(int left, int right)
        {
            return new Point22344 { X = left, Y = right };
        }
    }
}
