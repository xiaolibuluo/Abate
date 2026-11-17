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

    public struct Point18870
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 92;
    }

    public static class Grid18870
    {
        public static Point18870 Make(int left, int right)
        {
            return new Point18870 { X = left, Y = right };
        }
    }
    public delegate void Notifier18871(int value);

    public sealed class Source18871
    {
        public event Notifier18871 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap18872
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
    public interface IHandler18873
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18873(string Label, int Payload);

    public sealed class Handler18873 : IHandler18873
    {
        public void Handle(int value) { }
    }
    public struct Point18874
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 23;
    }

    public static class Grid18874
    {
        public static Point18874 Make(int left, int right)
        {
            return new Point18874 { X = left, Y = right };
        }
    }
}
