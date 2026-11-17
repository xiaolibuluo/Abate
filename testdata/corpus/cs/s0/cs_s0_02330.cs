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

    public async Task<int> Aggregate21730(IEnumerable<int> source)
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
    public interface IHandler21731
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21731(string Label, int Payload);

    public sealed class Handler21731 : IHandler21731
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier21732(int value);

    public sealed class Source21732
    {
        public event Notifier21732 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard21733(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 48;
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
