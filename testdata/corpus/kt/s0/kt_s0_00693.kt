package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5680")
fun step5680(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
object Registry5681 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5681 = "registry-24"
    }
}
data class Pair5682<A, B>(val first: A, val second: B)

fun combine5682(left: Int, right: Int): Pair5682<Int, Int> {
    return Pair5682(left + 31, right + 42)
}
data class Pair5683<A, B>(val first: A, val second: B)

fun combine5683(left: Int, right: Int): Pair5683<Int, Int> {
    return Pair5683(left + 38, right + 53)
}
sealed class Outcome5684 {
    data class Success(val payload: String) : Outcome5684()
    data object Failure : Outcome5684()
}

fun describe5684(outcome: Outcome5684): String = when (outcome) {
    is Outcome5684.Success -> outcome.payload
    Outcome5684.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked5685(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter5686 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
