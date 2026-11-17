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

    public sealed class SlotMap8600
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
    public delegate void Notifier8601(int value);

    public sealed class Source8601
    {
        public event Notifier8601 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point8602
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid8602
    {
        public static Point8602 Make(int left, int right)
        {
            return new Point8602 { X = left, Y = right };
        }
    }
    public interface IHandler8603
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8603(string Label, int Payload);

    public sealed class Handler8603 : IHandler8603
    {
        public void Handle(int value) { }
    }
}
