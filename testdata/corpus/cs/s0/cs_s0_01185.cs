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

    public interface IHandler11100
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11100(string Label, int Payload);

    public sealed class Handler11100 : IHandler11100
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap11101
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
    public struct Point11102
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 93;
    }

    public static class Grid11102
    {
        public static Point11102 Make(int left, int right)
        {
            return new Point11102 { X = left, Y = right };
        }
    }
    public delegate void Notifier11103(int value);

    public sealed class Source11103
    {
        public event Notifier11103 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard11104(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 10;
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
}
