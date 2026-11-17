package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry7280 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7280 = "registry-60"
    }
}
@JvmName("step7281")
fun step7281(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
@JvmName("step7282")
fun step7282(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
fun <T> List<T>.secondOrNull7283(): T? = if (size > 1) this[1] else null

fun render7283(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7284<A, B>(val first: A, val second: B)

fun combine7284(left: Int, right: Int): Pair7284<Int, Int> {
    return Pair7284(left + 88, right + 25)
}
data class Pair7285<A, B>(val first: A, val second: B)

fun combine7285(left: Int, right: Int): Pair7285<Int, Int> {
    return Pair7285(left + 95, right + 36)
}
object Registry7286 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7286 = "registry-5"
    }
}
sealed class Outcome7287 {
    data class Success(val payload: String) : Outcome7287()
    data object Failure : Outcome7287()
}

fun describe7287(outcome: Outcome7287): String = when (outcome) {
    is Outcome7287.Success -> outcome.payload
    Outcome7287.Failure -> "failed"
}
@JvmName("step7288")
fun step7288(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
class Counter7289 {
    var current: Int by Delegates.observable(26) { _, _, new -> println(new) }

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
