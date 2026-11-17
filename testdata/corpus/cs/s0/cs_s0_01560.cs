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

    public delegate void Notifier14590(int value);

    public sealed class Source14590
    {
        public event Notifier14590 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14591
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14591(string Label, int Payload);

    public sealed class Handler14591 : IHandler14591
    {
        public void Handle(int value) { }
    }
    public struct Point14592
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid14592
    {
        public static Point14592 Make(int left, int right)
        {
            return new Point14592 { X = left, Y = right };
        }
    }
    public static int Scale14593<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 68; }
        return number;
    }
    public sealed class SlotMap14594
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
    public static int Guard14595(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 92;
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
