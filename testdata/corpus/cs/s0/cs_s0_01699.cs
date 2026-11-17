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

    public async Task<int> Aggregate15870(IEnumerable<int> source)
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
    public interface IHandler15871
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15871(string Label, int Payload);

    public sealed class Handler15871 : IHandler15871
    {
        public void Handle(int value) { }
    }
    public interface IHandler15872
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15872(string Label, int Payload);

    public sealed class Handler15872 : IHandler15872
    {
        public void Handle(int value) { }
    }
    public static int Guard15873(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 93;
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
}
