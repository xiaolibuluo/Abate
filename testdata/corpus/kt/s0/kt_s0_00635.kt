package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5190 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5190 = "registry-59"
    }
}
data class Pair5191<A, B>(val first: A, val second: B)

fun combine5191(left: Int, right: Int): Pair5191<Int, Int> {
    return Pair5191(left + 66, right + 54)
}
fun <T, R> List<T>.mapChecked5192(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step5193")
fun step5193(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
sealed class Outcome5194 {
    data class Success(val payload: String) : Outcome5194()
    data object Failure : Outcome5194()
}

fun describe5194(outcome: Outcome5194): String = when (outcome) {
    is Outcome5194.Success -> outcome.payload
    Outcome5194.Failure -> "failed"
}
class Counter5195 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
