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

    public interface IHandler21240
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21240(string Label, int Payload);

    public sealed class Handler21240 : IHandler21240
    {
        public void Handle(int value) { }
    }
    public static int Guard21241(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 76;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public interface IHandler21242
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21242(string Label, int Payload);

    public sealed class Handler21242 : IHandler21242
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate21243(IEnumerable<int> source)
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
