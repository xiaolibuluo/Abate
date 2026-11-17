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

    public static int Guard11040(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 98;
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
    public interface IHandler11041
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11041(string Label, int Payload);

    public sealed class Handler11041 : IHandler11041
    {
        public void Handle(int value) { }
    }
    public struct Point11042
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid11042
    {
        public static Point11042 Make(int left, int right)
        {
            return new Point11042 { X = left, Y = right };
        }
    }
    public sealed class SlotMap11043
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
