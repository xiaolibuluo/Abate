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

    public interface IHandler20270
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20270(string Label, int Payload);

    public sealed class Handler20270 : IHandler20270
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap20271
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
    public interface IHandler20272
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20272(string Label, int Payload);

    public sealed class Handler20272 : IHandler20272
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap20273
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
