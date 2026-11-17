package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12530 {
    data class Success(val payload: String) : Outcome12530()
    data object Failure : Outcome12530()
}

fun describe12530(outcome: Outcome12530): String = when (outcome) {
    is Outcome12530.Success -> outcome.payload
    Outcome12530.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12531(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12532")
fun step12532(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
data class Pair12533<A, B>(val first: A, val second: B)

fun combine12533(left: Int, right: Int): Pair12533<Int, Int> {
    return Pair12533(left + 19, right + 29)
}
object Registry12534 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12534 = "registry-26"
    }
}
class Counter12535 {
    var current: Int by Delegates.observable(33) { _, _, new -> println(new) }

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
sealed class Outcome12536 {
    data class Success(val payload: String) : Outcome12536()
    data object Failure : Outcome12536()
}

fun describe12536(outcome: Outcome12536): String = when (outcome) {
    is Outcome12536.Success -> outcome.payload
    Outcome12536.Failure -> "failed"
}
