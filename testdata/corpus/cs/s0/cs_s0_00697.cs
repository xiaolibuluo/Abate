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

    public interface IHandler6520
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6520(string Label, int Payload);

    public sealed class Handler6520 : IHandler6520
    {
        public void Handle(int value) { }
    }
    public interface IHandler6521
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6521(string Label, int Payload);

    public sealed class Handler6521 : IHandler6521
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate6522(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
}
