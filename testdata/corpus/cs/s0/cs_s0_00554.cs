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

    public struct Point5190
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 59;
    }

    public static class Grid5190
    {
        public static Point5190 Make(int left, int right)
        {
            return new Point5190 { X = left, Y = right };
        }
    }
    public delegate void Notifier5191(int value);

    public sealed class Source5191
    {
        public event Notifier5191 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point5192
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 73;
    }

    public static class Grid5192
    {
        public static Point5192 Make(int left, int right)
        {
            return new Point5192 { X = left, Y = right };
        }
    }
    public interface IHandler5193
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5193(string Label, int Payload);

    public sealed class Handler5193 : IHandler5193
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap5194
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
