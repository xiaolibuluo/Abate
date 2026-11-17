package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10630<A, B>(val first: A, val second: B)

fun combine10630(left: Int, right: Int): Pair10630<Int, Int> {
    return Pair10630(left + 50, right + 69)
}
@JvmName("step10631")
fun step10631(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
fun <T> List<T>.secondOrNull10632(): T? = if (size > 1) this[1] else null

fun render10632(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10633<A, B>(val first: A, val second: B)

fun combine10633(left: Int, right: Int): Pair10633<Int, Int> {
    return Pair10633(left + 71, right + 41)
}
object Registry10634 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10634 = "registry-78"
    }
}
fun <T> List<T>.secondOrNull10635(): T? = if (size > 1) this[1] else null

fun render10635(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull10636(): T? = if (size > 1) this[1] else null

fun render10636(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step10637")
fun step10637(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
fun <T> List<T>.secondOrNull10638(): T? = if (size > 1) this[1] else null

fun render10638(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
