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

    public struct Point420
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 66;
    }

    public static class Grid420
    {
        public static Point420 Make(int left, int right)
        {
            return new Point420 { X = left, Y = right };
        }
    }
    public static int Scale421<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public sealed class SlotMap422
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
    public interface IHandler423
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope423(string Label, int Payload);

    public sealed class Handler423 : IHandler423
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier424(int value);

    public sealed class Source424
    {
        public event Notifier424 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
