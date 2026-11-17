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

    public delegate void Notifier12970(int value);

    public sealed class Source12970
    {
        public event Notifier12970 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler12971
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12971(string Label, int Payload);

    public sealed class Handler12971 : IHandler12971
    {
        public void Handle(int value) { }
    }
    public interface IHandler12972
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12972(string Label, int Payload);

    public sealed class Handler12972 : IHandler12972
    {
        public void Handle(int value) { }
    }
    public struct Point12973
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid12973
    {
        public static Point12973 Make(int left, int right)
        {
            return new Point12973 { X = left, Y = right };
        }
    }
    public sealed class SlotMap12974
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
