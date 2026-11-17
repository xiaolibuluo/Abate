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

    public static int Scale16440<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public delegate void Notifier16441(int value);

    public sealed class Source16441
    {
        public event Notifier16441 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier16442(int value);

    public sealed class Source16442
    {
        public event Notifier16442 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler16443
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16443(string Label, int Payload);

    public sealed class Handler16443 : IHandler16443
    {
        public void Handle(int value) { }
    }
    public struct Point16444
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 65;
    }

    public static class Grid16444
    {
        public static Point16444 Make(int left, int right)
        {
            return new Point16444 { X = left, Y = right };
        }
    }
    public sealed class SlotMap16445
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
