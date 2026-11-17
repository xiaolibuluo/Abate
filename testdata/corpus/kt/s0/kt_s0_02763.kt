package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22820<A, B>(val first: A, val second: B)

fun combine22820(left: Int, right: Int): Pair22820<Int, Int> {
    return Pair22820(left + 86, right + 62)
}
@JvmName("step22821")
fun step22821(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
class Counter22822 {
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
sealed class Outcome22823 {
    data class Success(val payload: String) : Outcome22823()
    data object Failure : Outcome22823()
}

fun describe22823(outcome: Outcome22823): String = when (outcome) {
    is Outcome22823.Success -> outcome.payload
    Outcome22823.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22824(): T? = if (size > 1) this[1] else null

fun render22824(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22825 {
    data class Success(val payload: String) : Outcome22825()
    data object Failure : Outcome22825()
}

fun describe22825(outcome: Outcome22825): String = when (outcome) {
    is Outcome22825.Success -> outcome.payload
    Outcome22825.Failure -> "failed"
}
object Registry22826 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22826 = "registry-31"
    }
}
fun <T> List<T>.secondOrNull22827(): T? = if (size > 1) this[1] else null

fun render22827(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
