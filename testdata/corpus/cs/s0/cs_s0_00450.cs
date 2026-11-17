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

    public sealed class SlotMap4220
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
    public interface IHandler4221
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4221(string Label, int Payload);

    public sealed class Handler4221 : IHandler4221
    {
        public void Handle(int value) { }
    }
    public struct Point4222
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 73;
    }

    public static class Grid4222
    {
        public static Point4222 Make(int left, int right)
        {
            return new Point4222 { X = left, Y = right };
        }
    }
    public delegate void Notifier4223(int value);

    public sealed class Source4223
    {
        public event Notifier4223 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
