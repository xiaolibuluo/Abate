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

    public delegate void Notifier22480(int value);

    public sealed class Source22480
    {
        public event Notifier22480 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22481<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
}
