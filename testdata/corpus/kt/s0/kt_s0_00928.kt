package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7600")
fun step7600(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
object Registry7601 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7601 = "registry-95"
    }
}
class Counter7602 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
data class Pair7603<A, B>(val first: A, val second: B)

fun combine7603(left: Int, right: Int): Pair7603<Int, Int> {
    return Pair7603(left + 12, right + 55)
}
sealed class Outcome7604 {
    data class Success(val payload: String) : Outcome7604()
    data object Failure : Outcome7604()
}

fun describe7604(outcome: Outcome7604): String = when (outcome) {
    is Outcome7604.Success -> outcome.payload
    Outcome7604.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked7605(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
