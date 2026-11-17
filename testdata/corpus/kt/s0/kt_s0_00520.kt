package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4210<A, B>(val first: A, val second: B)

fun combine4210(left: Int, right: Int): Pair4210<Int, Int> {
    return Pair4210(left + 46, right + 28)
}
@JvmName("step4211")
fun step4211(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 53
}
object Registry4212 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4212 = "registry-60"
    }
}
class Counter4213 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull4214(): T? = if (size > 1) this[1] else null

fun render4214(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome4215 {
    data class Success(val payload: String) : Outcome4215()
    data object Failure : Outcome4215()
}

fun describe4215(outcome: Outcome4215): String = when (outcome) {
    is Outcome4215.Success -> outcome.payload
    Outcome4215.Failure -> "failed"
}
