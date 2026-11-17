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

    public delegate void Notifier12880(int value);

    public sealed class Source12880
    {
        public event Notifier12880 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler12881
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12881(string Label, int Payload);

    public sealed class Handler12881 : IHandler12881
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap12882
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
