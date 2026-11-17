package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome1240 {
    data class Success(val payload: String) : Outcome1240()
    data object Failure : Outcome1240()
}

fun describe1240(outcome: Outcome1240): String = when (outcome) {
    is Outcome1240.Success -> outcome.payload
    Outcome1240.Failure -> "failed"
}
@JvmName("step1241")
fun step1241(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
object Registry1242 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1242 = "registry-79"
    }
}
fun <T> List<T>.secondOrNull1243(): T? = if (size > 1) this[1] else null

fun render1243(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair1244<A, B>(val first: A, val second: B)

fun combine1244(left: Int, right: Int): Pair1244<Int, Int> {
    return Pair1244(left + 93, right + 67)
}
class Counter1245 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
