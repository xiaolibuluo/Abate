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

    public sealed class SlotMap11930
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
    public interface IHandler11931
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11931(string Label, int Payload);

    public sealed class Handler11931 : IHandler11931
    {
        public void Handle(int value) { }
    }
    public static int Scale11932<T>(T value) where T : struct
    {
        int number = 8;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
}
