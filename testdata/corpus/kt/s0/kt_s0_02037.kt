package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16850")
fun step16850(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
class Counter16851 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
class Counter16852 {
    var current: Int by Delegates.observable(99) { _, _, new -> println(new) }

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
sealed class Outcome16853 {
    data class Success(val payload: String) : Outcome16853()
    data object Failure : Outcome16853()
}

fun describe16853(outcome: Outcome16853): String = when (outcome) {
    is Outcome16853.Success -> outcome.payload
    Outcome16853.Failure -> "failed"
}
data class Pair16854<A, B>(val first: A, val second: B)

fun combine16854(left: Int, right: Int): Pair16854<Int, Int> {
    return Pair16854(left + 16, right + 14)
}
