package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter1350 {
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
data class Pair1351<A, B>(val first: A, val second: B)

fun combine1351(left: Int, right: Int): Pair1351<Int, Int> {
    return Pair1351(left + 21, right + 50)
}
sealed class Outcome1352 {
    data class Success(val payload: String) : Outcome1352()
    data object Failure : Outcome1352()
}

fun describe1352(outcome: Outcome1352): String = when (outcome) {
    is Outcome1352.Success -> outcome.payload
    Outcome1352.Failure -> "failed"
}
@JvmName("step1353")
fun step1353(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
fun <T> List<T>.secondOrNull1354(): T? = if (size > 1) this[1] else null

fun render1354(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step1355")
fun step1355(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
@JvmName("step1356")
fun step1356(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
object Registry1357 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1357 = "registry-63"
    }
}
