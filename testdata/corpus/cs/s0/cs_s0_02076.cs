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

    public interface IHandler19360
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19360(string Label, int Payload);

    public sealed class Handler19360 : IHandler19360
    {
        public void Handle(int value) { }
    }
    public interface IHandler19361
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19361(string Label, int Payload);

    public sealed class Handler19361 : IHandler19361
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate19362(IEnumerable<int> source)
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
    public interface IHandler19363
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19363(string Label, int Payload);

    public sealed class Handler19363 : IHandler19363
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate19364(IEnumerable<int> source)
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
