package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step24500")
fun step24500(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
fun <T> List<T>.secondOrNull24501(): T? = if (size > 1) this[1] else null

fun render24501(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24502<A, B>(val first: A, val second: B)

fun combine24502(left: Int, right: Int): Pair24502<Int, Int> {
    return Pair24502(left + 53, right + 40)
}
data class Pair24503<A, B>(val first: A, val second: B)

fun combine24503(left: Int, right: Int): Pair24503<Int, Int> {
    return Pair24503(left + 60, right + 51)
}
object Registry24504 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24504 = "registry-67"
    }
}
