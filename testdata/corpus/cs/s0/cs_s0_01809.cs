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

    public static int Guard16880(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 27;
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
    public async Task<int> Aggregate16881(IEnumerable<int> source)
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
    public interface IHandler16882
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16882(string Label, int Payload);

    public sealed class Handler16882 : IHandler16882
    {
        public void Handle(int value) { }
    }
    public interface IHandler16883
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16883(string Label, int Payload);

    public sealed class Handler16883 : IHandler16883
    {
        public void Handle(int value) { }
    }
}
