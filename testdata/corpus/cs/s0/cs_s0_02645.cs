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

    public delegate void Notifier24750(int value);

    public sealed class Source24750
    {
        public event Notifier24750 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier24751(int value);

    public sealed class Source24751
    {
        public event Notifier24751 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap24752
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
    public delegate void Notifier24753(int value);

    public sealed class Source24753
    {
        public event Notifier24753 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point24754
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid24754
    {
        public static Point24754 Make(int left, int right)
        {
            return new Point24754 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate24755(IEnumerable<int> source)
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
