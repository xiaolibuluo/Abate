package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14820<A, B>(val first: A, val second: B)

fun combine14820(left: Int, right: Int): Pair14820<Int, Int> {
    return Pair14820(left + 65, right + 13)
}
@JvmName("step14821")
fun step14821(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
@JvmName("step14822")
fun step14822(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
fun <T> List<T>.secondOrNull14823(): T? = if (size > 1) this[1] else null

fun render14823(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14824 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14824 = "registry-93"
    }
}
sealed class Outcome14825 {
    data class Success(val payload: String) : Outcome14825()
    data object Failure : Outcome14825()
}

fun describe14825(outcome: Outcome14825): String = when (outcome) {
    is Outcome14825.Success -> outcome.payload
    Outcome14825.Failure -> "failed"
}
