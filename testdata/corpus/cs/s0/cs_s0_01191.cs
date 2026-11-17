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

    public sealed class SlotMap11160
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
    public delegate void Notifier11161(int value);

    public sealed class Source11161
    {
        public event Notifier11161 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier11162(int value);

    public sealed class Source11162
    {
        public event Notifier11162 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point11163
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid11163
    {
        public static Point11163 Make(int left, int right)
        {
            return new Point11163 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate11164(IEnumerable<int> source)
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
