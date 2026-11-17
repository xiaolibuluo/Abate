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

    public sealed class SlotMap26060
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
    public static int Scale26061<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public static int Scale26062<T>(T value) where T : struct
    {
        int number = 44;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public interface IHandler26063
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26063(string Label, int Payload);

    public sealed class Handler26063 : IHandler26063
    {
        public void Handle(int value) { }
    }
}
