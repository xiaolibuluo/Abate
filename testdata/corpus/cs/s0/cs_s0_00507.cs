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

    public sealed class SlotMap4750
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
    public static int Guard4751(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 76;
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
    public delegate void Notifier4752(int value);

    public sealed class Source4752
    {
        public event Notifier4752 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point4753
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 90;
    }

    public static class Grid4753
    {
        public static Point4753 Make(int left, int right)
        {
            return new Point4753 { X = left, Y = right };
        }
    }
}
