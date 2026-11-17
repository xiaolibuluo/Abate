package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17020 {
    data class Success(val payload: String) : Outcome17020()
    data object Failure : Outcome17020()
}

fun describe17020(outcome: Outcome17020): String = when (outcome) {
    is Outcome17020.Success -> outcome.payload
    Outcome17020.Failure -> "failed"
}
data class Pair17021<A, B>(val first: A, val second: B)

fun combine17021(left: Int, right: Int): Pair17021<Int, Int> {
    return Pair17021(left + 22, right + 39)
}
data class Pair17022<A, B>(val first: A, val second: B)

fun combine17022(left: Int, right: Int): Pair17022<Int, Int> {
    return Pair17022(left + 29, right + 50)
}
fun <T> List<T>.secondOrNull17023(): T? = if (size > 1) this[1] else null

fun render17023(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17024")
fun step17024(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
class Counter17025 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
sealed class Outcome17026 {
    data class Success(val payload: String) : Outcome17026()
    data object Failure : Outcome17026()
}

fun describe17026(outcome: Outcome17026): String = when (outcome) {
    is Outcome17026.Success -> outcome.payload
    Outcome17026.Failure -> "failed"
}
object Registry17027 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17027 = "registry-64"
    }
}
class Counter17028 {
    var current: Int by Delegates.observable(71) { _, _, new -> println(new) }

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
