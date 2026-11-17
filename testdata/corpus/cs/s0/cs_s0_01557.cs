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

    public static int Scale14570<T>(T value) where T : struct
    {
        int number = 31;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public delegate void Notifier14571(int value);

    public sealed class Source14571
    {
        public event Notifier14571 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14572
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14572(string Label, int Payload);

    public sealed class Handler14572 : IHandler14572
    {
        public void Handle(int value) { }
    }
    public struct Point14573
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 52;
    }

    public static class Grid14573
    {
        public static Point14573 Make(int left, int right)
        {
            return new Point14573 { X = left, Y = right };
        }
    }
    public sealed class SlotMap14574
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
    public struct Point14575
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 66;
    }

    public static class Grid14575
    {
        public static Point14575 Make(int left, int right)
        {
            return new Point14575 { X = left, Y = right };
        }
    }
}
