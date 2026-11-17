package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step20320")
fun step20320(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
fun <T> List<T>.secondOrNull20321(): T? = if (size > 1) this[1] else null

fun render20321(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20322")
fun step20322(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
data class Pair20323<A, B>(val first: A, val second: B)

fun combine20323(left: Int, right: Int): Pair20323<Int, Int> {
    return Pair20323(left + 58, right + 53)
}
@JvmName("step20324")
fun step20324(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
data class Pair20325<A, B>(val first: A, val second: B)

fun combine20325(left: Int, right: Int): Pair20325<Int, Int> {
    return Pair20325(left + 72, right + 14)
}
object Registry20326 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20326 = "registry-79"
    }
}
object Registry20327 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20327 = "registry-86"
    }
}
sealed class Outcome20328 {
    data class Success(val payload: String) : Outcome20328()
    data object Failure : Outcome20328()
}

fun describe20328(outcome: Outcome20328): String = when (outcome) {
    is Outcome20328.Success -> outcome.payload
    Outcome20328.Failure -> "failed"
}
