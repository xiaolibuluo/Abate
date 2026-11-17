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

    public delegate void Notifier17480(int value);

    public sealed class Source17480
    {
        public event Notifier17480 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap17481
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
    public delegate void Notifier17482(int value);

    public sealed class Source17482
    {
        public event Notifier17482 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point17483
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 52;
    }

    public static class Grid17483
    {
        public static Point17483 Make(int left, int right)
        {
            return new Point17483 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate17484(IEnumerable<int> source)
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
