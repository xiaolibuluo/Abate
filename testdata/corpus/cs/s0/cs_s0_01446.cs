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

    public interface IHandler13550
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13550(string Label, int Payload);

    public sealed class Handler13550 : IHandler13550
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier13551(int value);

    public sealed class Source13551
    {
        public event Notifier13551 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point13552
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 77;
    }

    public static class Grid13552
    {
        public static Point13552 Make(int left, int right)
        {
            return new Point13552 { X = left, Y = right };
        }
    }
    public sealed class SlotMap13553
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
    public struct Point13554
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 91;
    }

    public static class Grid13554
    {
        public static Point13554 Make(int left, int right)
        {
            return new Point13554 { X = left, Y = right };
        }
    }
}
