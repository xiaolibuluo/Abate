package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter4350 {
    var current: Int by Delegates.observable(34) { _, _, new -> println(new) }

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
@JvmName("step4351")
fun step4351(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
sealed class Outcome4352 {
    data class Success(val payload: String) : Outcome4352()
    data object Failure : Outcome4352()
}

fun describe4352(outcome: Outcome4352): String = when (outcome) {
    is Outcome4352.Success -> outcome.payload
    Outcome4352.Failure -> "failed"
}
object Registry4353 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4353 = "registry-55"
    }
}
object Registry4354 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4354 = "registry-62"
    }
}
sealed class Outcome4355 {
    data class Success(val payload: String) : Outcome4355()
    data object Failure : Outcome4355()
}

fun describe4355(outcome: Outcome4355): String = when (outcome) {
    is Outcome4355.Success -> outcome.payload
    Outcome4355.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4356(): T? = if (size > 1) this[1] else null

fun render4356(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4357<A, B>(val first: A, val second: B)

fun combine4357(left: Int, right: Int): Pair4357<Int, Int> {
    return Pair4357(left + 83, right + 20)
}
