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

    public static int Guard5690(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 30;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public struct Point5691
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid5691
    {
        public static Point5691 Make(int left, int right)
        {
            return new Point5691 { X = left, Y = right };
        }
    }
    public sealed class SlotMap5692
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
    public delegate void Notifier5693(int value);

    public sealed class Source5693
    {
        public event Notifier5693 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
