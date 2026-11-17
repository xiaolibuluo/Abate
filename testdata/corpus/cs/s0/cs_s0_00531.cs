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

    public interface IHandler4970
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4970(string Label, int Payload);

    public sealed class Handler4970 : IHandler4970
    {
        public void Handle(int value) { }
    }
    public interface IHandler4971
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4971(string Label, int Payload);

    public sealed class Handler4971 : IHandler4971
    {
        public void Handle(int value) { }
    }
}
