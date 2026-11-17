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

    public sealed class SlotMap1740
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
    public delegate void Notifier1741(int value);

    public sealed class Source1741
    {
        public event Notifier1741 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler1742
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1742(string Label, int Payload);

    public sealed class Handler1742 : IHandler1742
    {
        public void Handle(int value) { }
    }
    public static int Scale1743<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 40; }
        return number;
    }
}
