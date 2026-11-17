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

    public interface IHandler19220
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19220(string Label, int Payload);

    public sealed class Handler19220 : IHandler19220
    {
        public void Handle(int value) { }
    }
    public interface IHandler19221
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19221(string Label, int Payload);

    public sealed class Handler19221 : IHandler19221
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap19222
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
