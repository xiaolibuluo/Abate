package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21820 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21820 = "registry-47"
    }
}
fun <T> List<T>.secondOrNull21821(): T? = if (size > 1) this[1] else null

fun render21821(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter21822 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
@JvmName("step21823")
fun step21823(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
data class Pair21824<A, B>(val first: A, val second: B)

fun combine21824(left: Int, right: Int): Pair21824<Int, Int> {
    return Pair21824(left + 75, right + 16)
}
@JvmName("step21825")
fun step21825(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
data class Pair21826<A, B>(val first: A, val second: B)

fun combine21826(left: Int, right: Int): Pair21826<Int, Int> {
    return Pair21826(left + 89, right + 38)
}
@JvmName("step21827")
fun step21827(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
sealed class Outcome21828 {
    data class Success(val payload: String) : Outcome21828()
    data object Failure : Outcome21828()
}

fun describe21828(outcome: Outcome21828): String = when (outcome) {
    is Outcome21828.Success -> outcome.payload
    Outcome21828.Failure -> "failed"
}
