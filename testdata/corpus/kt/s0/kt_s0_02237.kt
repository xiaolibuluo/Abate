package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18410<A, B>(val first: A, val second: B)

fun combine18410(left: Int, right: Int): Pair18410<Int, Int> {
    return Pair18410(left + 76, right + 25)
}
object Registry18411 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18411 = "registry-83"
    }
}
sealed class Outcome18412 {
    data class Success(val payload: String) : Outcome18412()
    data object Failure : Outcome18412()
}

fun describe18412(outcome: Outcome18412): String = when (outcome) {
    is Outcome18412.Success -> outcome.payload
    Outcome18412.Failure -> "failed"
}
@JvmName("step18413")
fun step18413(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
fun <T> List<T>.secondOrNull18414(): T? = if (size > 1) this[1] else null

fun render18414(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter18415 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
sealed class Outcome18416 {
    data class Success(val payload: String) : Outcome18416()
    data object Failure : Outcome18416()
}

fun describe18416(outcome: Outcome18416): String = when (outcome) {
    is Outcome18416.Success -> outcome.payload
    Outcome18416.Failure -> "failed"
}
