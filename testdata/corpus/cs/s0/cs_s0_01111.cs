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

    public interface IHandler10410
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10410(string Label, int Payload);

    public sealed class Handler10410 : IHandler10410
    {
        public void Handle(int value) { }
    }
    public interface IHandler10411
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10411(string Label, int Payload);

    public sealed class Handler10411 : IHandler10411
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap10412
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
