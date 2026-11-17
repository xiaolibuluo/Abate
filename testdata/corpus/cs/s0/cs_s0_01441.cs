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

    public interface IHandler13500
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13500(string Label, int Payload);

    public sealed class Handler13500 : IHandler13500
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate13501(IEnumerable<int> source)
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
    public interface IHandler13502
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13502(string Label, int Payload);

    public sealed class Handler13502 : IHandler13502
    {
        public void Handle(int value) { }
    }
    public static int Guard13503(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 19;
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
