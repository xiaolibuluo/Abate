package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18460 {
    var current: Int by Delegates.observable(44) { _, _, new -> println(new) }

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
data class Pair18461<A, B>(val first: A, val second: B)

fun combine18461(left: Int, right: Int): Pair18461<Int, Int> {
    return Pair18461(left + 51, right + 10)
}
sealed class Outcome18462 {
    data class Success(val payload: String) : Outcome18462()
    data object Failure : Outcome18462()
}

fun describe18462(outcome: Outcome18462): String = when (outcome) {
    is Outcome18462.Success -> outcome.payload
    Outcome18462.Failure -> "failed"
}
@JvmName("step18463")
fun step18463(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
object Registry18464 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18464 = "registry-72"
    }
}
@JvmName("step18465")
fun step18465(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
object Registry18466 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18466 = "registry-86"
    }
}
@JvmName("step18467")
fun step18467(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
sealed class Outcome18468 {
    data class Success(val payload: String) : Outcome18468()
    data object Failure : Outcome18468()
}

fun describe18468(outcome: Outcome18468): String = when (outcome) {
    is Outcome18468.Success -> outcome.payload
    Outcome18468.Failure -> "failed"
}
